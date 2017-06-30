<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="initial-scale=1">
    <title>ECharts</title>
    <!-- 引入 echarts.js -->
    <script src="/static/js/echarts.min.js"></script>
    <script src="/static/js/jquery.min.js"></script>
</head>
<body>
    <!-- 为ECharts准备一个具备大小（宽高）的Dom -->
    <div id="main" style="width: 100%; height:400px"></div>
    <script type="text/javascript">
        // 基于准备好的dom，初始化echarts实例
        var myChart = echarts.init(document.getElementById('main'));

        var dimension = {{.Dimension}}

        var style = {
                   normal: {
                     label : {
                         show : true,
                         position : 'top',
                         textStyle : {
                             fontSize : '20',
                             fontFamily : '微软雅黑',
                             fontWeight : 'bold',
                             color: 'blue'
                         }
                      }
                    }
                 }

        // 指定图表的配置项和数据
        var option = {
            title: {
                text: '{{.Title}}',
                x: 'center'
            },
            toolbox: {
                show: true,
                x: 'right',
                orient:'vertical',
                feature: {
                    dataView: {readOnly: false},
                    restore: {},
                    saveAsImage: {}
                }
            },
            legend: {
                selectedMode:'multiple',
                x: 'left',
                orient:'vertical',
                data: {{.Dimension}}
            },
            xAxis: [{
               axisLabel : {
                 interval : 0,
                 rotate : 90
               },
               data:{{.XToday}}
            }],
            yAxis: {
              type : 'value'
            },
            series: [{
                name: dimension[0],
                type: 'bar',
                data: {{.YToday}},
                xAxisIndex: 0,
                itemStyle: style
            },  {
                name: dimension[1],
                type: 'bar',
                data: {{.YYestoday}},
                xAxisIndex: 0,
                itemStyle: style
            }]
        };

        // 使用刚指定的配置项和数据显示图表。
        myChart.setOption(option);
        $(window).on('resize', function(){
        if(myChart != null && myChart != undefined){
            myChart.resize();
        };
    });
    </script>
</body>
</html>
