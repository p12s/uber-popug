import {applyMiddleware, combineReducers, createStore} from 'redux';
import {composeWithDevTools} from 'redux-devtools-extension';
import thunk from 'redux-thunk';
import accountReducer from './accountReducer';


const rootReducer = combineReducers({
  account: accountReducer,
})

export const store = createStore(rootReducer, 
  composeWithDevTools(applyMiddleware(thunk)))
