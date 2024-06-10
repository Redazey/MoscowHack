import React from 'react'
import Header from "./components/header/Header.tsx";
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Footer from "./components/footer/Footer.tsx";
import NewsBlock from "./components/newsBlock/NewsBlock.tsx";


function App() {

  return (
      <Router>
          <Header/>
            <NewsBlock/>
          <Footer/>
      </Router>
  )
}

export default App
