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

        var dimension = {{.YName}}
        var mydata = {{.YData}}

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
            tooltip : {
                  trigger: 'axis',
                  axisPointer: {
                      type: 'cross',
                      label: {
                          backgroundColor: '#6a7985'
                      }
                  }
            },
            dataZoom: [{
                startValue: {{.Start}}
            }, {
                type: 'inside'
            }],
            toolbox: {
                show: true,
                x: 'left',
                orient:'vertical',
                feature: {
                   saveAsImage: {}
                }
            },
            legend: {
                x: 'right',
                orient:'vertical',
                data: {{.YName}}
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            xAxis: [{
               boundaryGap : false,
               data:{{.X}},
               splitLine: {
                   show: false
               }
            }],
            yAxis: {
                type: 'value'
            },
            series: [{
                name: dimension[0],
                type: 'line',
                stack: '总量',
                label: {
                    normal: {
                        show: true,
                        position: 'top'
                    }
                },
                data: mydata[0],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[1],
                type: 'line',
                stack: '总量',
                data: mydata[1],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[2],
                type: 'line',
                stack: '总量',
                data: mydata[2],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[3],
                type: 'line',
                stack: '总量',
                data: mydata[3],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[4],
                type: 'line',
                stack: '总量',
                data: mydata[4],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[5],
                type: 'line',
                stack: '总量',
                data: mydata[5],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[6],
                type: 'line',
                stack: '总量',
                data: mydata[6],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            },  {
                name: dimension[7],
                type: 'line',
                stack: '总量',
                data: mydata[7],
                areaStyle: {normal: {}},
                xAxisIndex: 0,
                showSymbol: false,
                hoverAnimation: false,
                itemStyle: style
            }]
        };

        // 使用刚指定的配置项和数据显示图表。
        myChart.setOption(option);
    </script>
</body>
</html>
