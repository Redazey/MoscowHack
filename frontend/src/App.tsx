import React from 'react'
import Header from "./components/header/Header.tsx";
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Footer from "./components/footer/Footer.tsx";


function App() {

  return (
      <Router>
          <Header/>
          <Footer/>
      </Router>
  )
}

export default App
