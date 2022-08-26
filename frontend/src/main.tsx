import React from 'react';
import { createRoot } from 'react-dom/client';
import { HashRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import 'antd/dist/antd.css';
import './style.css';
import App from './App';
import store from './store';

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <React.StrictMode>
      <Provider store={store}>
        <HashRouter basename="/">
          <App />
        </HashRouter>
      </Provider>
    </React.StrictMode>
  );
} else {
  throw new Error("can't find the root element");
}
