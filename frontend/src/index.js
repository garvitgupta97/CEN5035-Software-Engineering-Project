import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

import Home from "./Home/Home";

import { BrowserRouter } from "react-router-dom";
import * as serviceWorker from "./serviceWorker";

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

ReactDOM.render(
  <BrowserRouter>
    
    <Switch>
        <App />
        <Route path="/home" component={Home} />
      </Switch>
    </BrowserRouter>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
serviceWorker.unregister();




