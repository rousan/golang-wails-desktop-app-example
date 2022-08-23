import React from 'react';
import { createRoot } from 'react-dom/client';
import { HashRouter } from 'react-router-dom';
import './style.css';
import App from './App';

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(
    <React.StrictMode>
      <HashRouter basename="/">
        <App />
      </HashRouter>
    </React.StrictMode>
  );
} else {
  throw new Error("can't find the root element");
}

