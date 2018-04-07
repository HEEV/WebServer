function renderGraph(runNumber) {
  $.post('./API/GetGraph.php', { runId: runNumber }, function(data) {
    var arr = JSON.parse(data);
    console.log(arr);
    var dataArray = [];
    $.each(arr, function(index, value) {
      if (index !== 'LogTime' && index !== 'RunNumber'
          && index !== 'Id' && index !== 'CarId') {
        dataArray.push({ 
          name: index,
          data: value,
          visible: ((index === 'Id' ||
                     index === 'CarId' ||
                     index === 'Latitude' ||
                     index === 'Longitude'||
                     index === 'LKillSwitch' ||
                     index === 'MKillSwitch' ||
                     index === 'RKillSwitch') ? false : true)
        });
      }
    });
    var timeArray = [];
    $.each(arr.LogTime, function(index, value) {
      timeArray.push(value.substring(11));
    });
    Highcharts.chart('container', {
      title: {
        text: 'Cedarville Supermileage run data for '
        + $('#RunNum :selected').text().split(' ')[0]
        + ' on ' + arr.LogTime[0].substring(0, 10)
      },

      xAxis: {
        categories: timeArray
      },

      yAxis: {
        title: {
          text: ''
        }
      },

      legend: {
        layout: 'vertical',
        align: 'right',
        verticalAlign: 'middle'
      },

      plotOptions: {
        series: {
          label: {
            connectorAllowed: false
          },
        }
      },

      series: dataArray,
      responsive: {
        rules: [{
          condition: {
            maxWidth: 500
          },
          chartOptions: {
            legend: {
              layout: 'horizontal',
              align: 'center',
              verticalAlign: 'bottom'
            }
          }
        }]
      }
    });
  });
}

$(function() {
  $.get('./API/GetRunIds.php', function(data) {
    var arr = JSON.parse(data);
    console.log(arr);
    $.each(arr, function(index, value) {
      $('#RunNum').append(
        '<option value="' + value.RunId + '">' + value.Car + ' on ' +
        value.Time.substring(0, 19) + '</option>'
      );
    });

    renderGraph($('#RunNum').val());
  });


  $('#RunNum').on('change', function() {
    renderGraph($(this).val());
  });
});
