import React from 'react'
import Header from "./components/header/Header.tsx";
import { BrowserRouter, Route, Routes } from 'react-router-dom';


function App() {

  return (
      <>
          <BrowserRouter>
              <div>
                  <main>
                      <Routes>
                          <Route path="/" element={<Header />} />
                          {/* <Route path="/id/:id" element={<Movie />} /> */}
                      </Routes>
                  </main>
              </div>
          </BrowserRouter>
      </>
  )
}

export default App
