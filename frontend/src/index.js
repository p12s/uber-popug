import React from 'react';
import ReactDOM from 'react-dom';
import history from "./history";
import App from './App';
import en_US from '@douyinfe/semi-ui/lib/es/locale/source/en_US';
import { LocaleProvider } from '@douyinfe/semi-ui';
import {store} from './reducers';
import {Provider} from 'react-redux';

ReactDOM.render(
  <Provider store={store} history={history}>
    <LocaleProvider locale={en_US}>
      <App />
    </LocaleProvider>
  </Provider>,  
  document.getElementById('root')
);
