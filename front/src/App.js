import React from 'react';
import { Provider } from "react-redux";
import { BrowserRouter as Router, Route } from "react-router-dom";

import Login from "./components/Login"
import ChatList from "./components/ChatList"
import MessageBlock from "./components/MessageBlock"

import {store} from "./store/stora"
import './App.css';

const App = () => (
    <Provider store={store}>
        <Router>
            <div>
                <Route exact path="/" component={Login} />
                {/*<Route path='/messages' component={ChatList}/>*/}
                {/*<Route path='/messages/:id' component={MessageBlock}/>*/}
            </div>
        </Router>
    </Provider>
);

export default App;
