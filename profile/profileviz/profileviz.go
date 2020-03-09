package profileviz

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/jamesrr39/goutil/errorsx"
	"github.com/jamesrr39/goutil/profile"
	"github.com/jamesrr39/goutil/streamtostorage"
)

func Generate(dataFilePath, outFilePath string) errorsx.Error {
	file, err := os.Open(dataFilePath)
	if err != nil {
		return errorsx.Wrap(err)
	}
	defer file.Close()

	runs := []profile.Run{}

	reader := streamtostorage.NewReader(file)
	for {
		b, err := reader.ReadNextMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			return errorsx.Wrap(err)
		}

		var run profile.Run
		err = proto.Unmarshal(b, &run)
		if err != nil {
			return errorsx.Wrap(err)
		}

		runs = append(runs, run)
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return errorsx.Wrap(err)
	}
	defer outFile.Close()

	runsBytes, err := json.Marshal(runs)
	if err != nil {
		return errorsx.Wrap(err)
	}

	_, err = fmt.Fprintf(outFile, tpl, chartJSCSS, chartJSJS, string(runsBytes))
	if err != nil {
		return errorsx.Wrap(err)
	}

	return nil
}

const tpl = `<html>
	<head>
		<meta content="text/html;charset=utf-8" http-equiv="Content-Type">
		<meta content="utf-8" http-equiv="encoding">
		<title>Profile</title>
		<style type="text/css">
			%s
		</style>
		<script>
			%s
		</script>
		<script>
			var runs = %s;
		</script>
		<script>
			function generateColor() {
				const colors = [...Array(3)].map(() => Math.floor(Math.random() * 256));

				return "rgb(" + colors.join(', ') + ")";
			}

			function drawCanvas(run) {
				const runEvents = run.events || [];
				const eventsDatasets = runEvents.map(event => {
					const c = generateColor();
					return {
						label: event.name,
						data: [event.timeNanos / (1000 * 1000)],
						// fill: false,
						borderColor: c,
					};
				});

				const maxItemsToDisplay = 200;
				if (eventsDatasets.length > maxItemsToDisplay) {
					alert("too many items to display (max " + maxItemsToDisplay + ", got " + eventsDatasets.length + ")");
					return;
				}

				const datasets = [{
					label: "Start",
					data: [run.startTimeNanos / (1000 * 1000)],
					fill: false,
					backgroundColor: 'blue',
				}].concat(eventsDatasets).concat([{
					label: "End",
					data: [run.endTimeNanos / (1000 * 1000)],
					fill: false,
					backgroundColor: 'blue',
				}]);

				var ctx = document.getElementById("canvas").getContext("2d");
				var data = {
					labels: ["Events"],
					datasets,
				};
				var options = {
					// Elements options apply to all of the options unless overridden in a dataset
					// In this case, we are setting the border of each horizontal bar to be 2px wide
					elements: {
						rectangle: {
							borderWidth: 2,
						}
					},
					responsive: true,
					legend: {
						position: 'right',
					},
					title: {
						display: true,
						text: run.name + (run.summary && " (" + run.summary + ")"),
					},
					barBeginAtOrigin: true,
					// scales: {
						// xAxes: [{
						// 	stacked: true
						// }],
					// }
				};
				window.myBarChart && window.myBarChart.destroy();
				window.myBarChart = new Chart(ctx, {
					type: 'horizontalBar',
					// type: 'line',
					data: data,
					options: options
				});
			}
		</script>
	</head>
	<body>
		Chart
		<div id="table-wrapper"></div>
		<canvas id="canvas"></canvas>
		<script>
			const tableKeys = ['name', 'summary', 'startTimeNanos', 'endTimeNanos'];
			const table = document.createElement('table');
			const thead = document.createElement('thead');
			table.appendChild(thead);
			const theadTr = document.createElement('tr');
			thead.appendChild(theadTr);

			tableKeys.concat(['duration milliseconds', 'draw graph']).forEach(key => {
				const th = document.createElement('th');
				th.innerText = key;
				theadTr.appendChild(th);
			});

			const tbody = document.createElement('tbody');
			table.appendChild(tbody);

			const tableRows = runs.map(run => {
				const tr = document.createElement('tr');
				tableKeys.forEach(key => {
					const td = document.createElement('td');
					td.innerText = run[key];
					tr.appendChild(td);
				});

				const durationTd = document.createElement('td');
				durationTd.innerText = (run.endTimeNanos - run.startTimeNanos) / 10**6
				tr.appendChild(durationTd);

				const drawGraphBtn = document.createElement('button');
				drawGraphBtn.innerText = 'draw graph';
				drawGraphBtn.addEventListener('click', () => {
					drawCanvas(run);
				});

				const drawGraphTd = document.createElement('td');
				drawGraphTd.appendChild(drawGraphBtn);
				tr.appendChild(drawGraphTd);

				tbody.appendChild(tr);
			});
			document.getElementById('table-wrapper').innerHTML = '';
			document.getElementById('table-wrapper').appendChild(table);
		</script>
	</body>
</html>
`
