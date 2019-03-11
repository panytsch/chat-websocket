import { createStore, applyMiddleware, combineReducers } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";

function userData(
    state = {
        user: {
            name: null,
            token: null,
        },
        currentConnect: null,
    },
    action
) {
    switch (action.type) {
        case 'AUTH':
            console.log(action);
        default:
            return state;
    }
}

const reducers = combineReducers({ userData });

export const store = createStore(
    reducers,
    composeWithDevTools(applyMiddleware(thunk))
);