import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import Footer from "./components/footer/Footer.tsx";
import Header from "./components/header/Header.tsx";

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <App/>
  </React.StrictMode>,
)
