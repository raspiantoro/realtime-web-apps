import React, { Component } from "react";
import {Line} from 'react-chartjs-2';
import EventEmitter from "../../utils/eventemitter"
import Stream from "../../services/stream"

class Graph extends Component {
  

  constructor(props){
    super(props)

    this.state = {
      option:{
        animation: {
          duration: 0
        },
        maintainAspectRatio:false,
      },
      data: {
        labels: [1,2,3,4,5,6,7,8,9,10],
        datasets: [
          {
            label: 'Stream Data',
            fill: true,
            lineTension: 0.1,
            backgroundColor: 'rgba(75,192,192,0.4)',
            borderColor: 'rgba(75,192,192,1)',
            borderCapStyle: 'butt',
            borderDash: [],
            borderDashOffset: 0.0,
            borderJoinStyle: 'miter',
            pointBorderColor: 'rgba(75,192,192,1)',
            pointBackgroundColor: '#fff',
            pointBorderWidth: 1,
            pointHoverRadius: 5,
            pointHoverBackgroundColor: 'rgba(75,192,192,1)',
            pointHoverBorderColor: 'rgba(220,220,220,1)',
            pointHoverBorderWidth: 2,
            pointRadius: 1,
            pointHitRadius: 10,
            data: [0,0,0,0,0,0,0,0,0,0]
          }
        ]
      }
    }
  }
  

  getStream(){
    const emitter = new EventEmitter()
    const stream = new Stream(emitter)
    
    stream.get()
    const updateChart = this.updateChart
    const parent = this

    emitter.on('newstream', function(data) {
      updateChart(data['count'], parent)
    })
  }  

  updateChart(data, parent){
    const datasetsCopy = parent.state.data.datasets.slice(0);
    
    const labelCopy = parent.state.data.labels
    let maxVal = labelCopy[labelCopy.length - 1]
    labelCopy.shift()
    labelCopy.push(maxVal + 1)

    const dataCopy = datasetsCopy[0].data.slice(0)
    dataCopy.shift()
    dataCopy.push(data)  
    datasetsCopy[0].data = dataCopy

    parent.setState({
        data: Object.assign({}, parent.state.data, {
            label: labelCopy,
            datasets: datasetsCopy
        })
    });
  }
  
  componentDidMount() {
    this.getStream()
  }

  render() {
    return (
      <div>
        <Line data={this.state.data} options={this.state.option} height={600}/>
      </div>
    );
  }
}
 
export default Graph;