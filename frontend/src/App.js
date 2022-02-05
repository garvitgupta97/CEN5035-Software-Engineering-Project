
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';

import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Login from "./login/login";
import SignUp from "./signup/signup";


import Home from './Home/Home'
import AboutUs from './AboutUs/AboutUs'
import Account from './Account/Account'


function App() {
  return (<Router>
    <div className="App">
   

          <Switch>
            <Route exact path='/' component={Login} />
            <Route path="/sign-in" component={Login} />
            <Route path="/sign-up" component={SignUp} />
            <Route path={'/AboutUs'} component={AboutUs}></Route>
            <Route path={'/Account'} component={Account}></Route>
            <Route path={'/home'} component={Home}></Route>
          </Switch>
    
    </div>


 </Router>
  );
}


export default App;
