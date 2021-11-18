import {applyMiddleware, combineReducers, createStore} from 'redux';
import {composeWithDevTools} from 'redux-devtools-extension';
import thunk from 'redux-thunk';
import accountReducer from './accountReducer';
import taskReducer from './taskReducer';


const rootReducer = combineReducers({
  account: accountReducer,
  task: taskReducer,
})

export const store = createStore(rootReducer, 
  composeWithDevTools(applyMiddleware(thunk)))
