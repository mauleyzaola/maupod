import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reducer from './reducers';
import { Provider } from 'react-redux';
import { createStore } from "redux";
import middleware from './middleware';

const store = createStore(reducer, middleware);

ReactDOM.render(
  <React.StrictMode>
      <Provider store={store}>
          <App />
      </Provider>
  </React.StrictMode>,
  document.getElementById('root')
);
