import React, { Component } from "react";
import { Route, NavLink, HashRouter } from "react-router-dom";
import Home from "./scenes/home";
import About from "./scenes/about";
 
class Main extends Component {
  render() {
    return (
      <HashRouter>
        <div>
          <h1 className="title">Monitoring</h1>
          <ul className="header">
            <li><NavLink to="/">Home</NavLink></li>
            <li><NavLink to="/about">About</NavLink></li>
          </ul>
          <div className="content">
            <Route exact path="/" component={Home}/>
            <Route path="/about" component={About}/>
          </div>
        </div>
      </HashRouter>
    );
  }
}
 
export default Main;