import React from 'react'
import Header from "./components/header/Header.tsx";
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Footer from "./components/footer/Footer.tsx";
import NewsBlock from "./components/newsBlock/NewsBlock.tsx";
import NewsSlider from "./components/newsBlock/test/NewsSlider.tsx";



function App() {

  return (
    <React.StrictMode>
      <Router>
          <Header/>
            <NewsBlock/>
          <Footer/>
      </Router>
    </React.StrictMode>
  )
}

export default App
