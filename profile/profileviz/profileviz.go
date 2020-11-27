package profileviz

import (
	"html/template"
	"io"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/jamesrr39/goutil/errorsx"
	"github.com/jamesrr39/goutil/profile"
	"github.com/jamesrr39/goutil/streamtostorage"
)

func streamToStorageReaderToRuns(reader *streamtostorage.Reader) ([]*RunType, errorsx.Error) {
	runs := []*RunType{}
	for {
		b, err := reader.ReadNextMessage()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errorsx.Wrap(err)
		}

		run := new(profile.Run)
		err = proto.Unmarshal(b, run)
		if err != nil {
			return nil, errorsx.Wrap(err)
		}

		runDuration := time.Duration(run.EndTimeNanos - run.StartTimeNanos)

		var events []*EventType
		for _, event := range run.Events {
			ratioThrough := float64(event.TimeNanos-run.StartTimeNanos) / float64(runDuration)
			events = append(events, &EventType{
				Name:                event.Name,
				TimeSinceStartOfRun: time.Duration(event.TimeNanos - run.StartTimeNanos),
				PercentageThrough:   ratioThrough * 100,
			})
		}

		summary := run.Summary
		if summary == "" {
			summary = "(none)"
		}
		startTimeSeconds := run.StartTimeNanos / (1000 * 1000 * 1000)

		vizRun := &RunType{
			Name:      run.Name,
			Summary:   summary,
			StartTime: time.Unix(startTimeSeconds, run.StartTimeNanos/startTimeSeconds).Format("2006-01-02T15:04:05.999"),
			Duration:  runDuration.String(),
			Events:    events,
		}

		runs = append(runs, vizRun)
	}

	return runs, nil
}

func Generate(dataFilePath, outFilePath string) errorsx.Error {
	file, err := os.Open(dataFilePath)
	if err != nil {
		return errorsx.Wrap(err)
	}
	defer file.Close()

	reader := streamtostorage.NewReader(file, streamtostorage.MessageSizeBufferLenDefault)

	runs, err := streamToStorageReaderToRuns(reader)
	if err != nil {
		return errorsx.Wrap(err)
	}

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return errorsx.Wrap(err)
	}
	defer outFile.Close()

	data := &tmplDataType{
		Runs: runs,
	}

	err = gotpl.Execute(outFile, data)
	if err != nil {
		return errorsx.Wrap(err)
	}
	// runsBytes, err := json.Marshal(runs)
	// if err != nil {
	// 	return errorsx.Wrap(err)
	// }

	// _, err = fmt.Fprintf(outFile, tpl, chartJSCSS, chartJSJS, string(runsBytes))
	// if err != nil {
	// 	return errorsx.Wrap(err)
	// }

	return nil
}

type EventType struct {
	Name                string
	PercentageThrough   float64
	TimeSinceStartOfRun time.Duration
}

type RunType struct {
	Name, Summary, StartTime, Duration string
	Events                             []*EventType
}

type tmplDataType struct {
	Runs []*RunType
}

var gotpl = template.Must(template.New("profileviz").Parse(`
	<html>
		<head>
			<meta content="text/html;charset=utf-8" http-equiv="Content-Type">
			<meta content="utf-8" http-equiv="encoding">
			<title>Profile</title>
			<style type="text/css">
				.events-table {
					width: 100%;
					background: lightblue;
				}
				.event-percentage-through-cell {
					width: 100%;
					border-left: 1px solid grey;
					border-right: 1px solid grey;
				}
				.event-name {
					min-width: 100px;
				}
				.event-since-start-of-run {
					min-width: 100px;
				}
			</style>
		</head>
		<body>
			<table cellspacing="10px">
				<thead>
					<tr>
						<th></th>
						<th>Run</th>
						<th>Summary</th>
						<th>Start Time</th>
						<th>Duration</th>
					</tr>
				</thead>
				<tbody>
				{{range .Runs}}
					<tr>
						<td>
							<button type="button" class="expand-events-row">Expand</button>
						</td>
						<td>{{.Name}}</td>
						<td>{{.Summary}}</td>
						<td>{{.StartTime}}</td>
						<td>{{.Duration}}</td>
					</tr>
					<tr style="display: none;">
						<td colspan="5">
							<table class="events-table">
								<tbody>
									{{range .Events}}
										<tr>
											<td class="event-name">{{.Name}}</td>
											<td class="event-since-start-of-run">{{.TimeSinceStartOfRun}}</td>
											<td title="{{.PercentageThrough}}%" class="event-percentage-through-cell">
												<span style="border-left: 1px solid blue; margin-left:{{.PercentageThrough}}%"></span>
											</td>
										</tr>
									{{end}}
								</tbody>
							</table>
						</td>
					</tr>
				{{end}}
				</tbody>
			</table>
			<script>
				for (let expandButton of document.getElementsByClassName("expand-events-row")) {
					expandButton.addEventListener("click", event => {
						console.log(event);
						const detailRow = event.target.parentElement.parentElement.nextElementSibling;
						if (detailRow.style.display === "none") {
							detailRow.style.display = "table-row";
							expandButton.innerText = "Collapse";
						} else {
							detailRow.style.display = "none";
							expandButton.innerText = "Expand";
						}
					});
				}
			</script>
		</body>
	</html>
`))

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
