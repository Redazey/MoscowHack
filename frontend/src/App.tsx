import React from 'react'
import Header from "./components/header/Header.tsx";
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import Footer from "./components/footer/Footer.tsx";
import NewsBlock from "./components/newsBlock/NewsBlock.tsx";
import VacanciesBlock from "./components/VacanciesBlock/VacancyBlock.tsx";
import VacancyBlock from "./components/VacanciesBlock/VacancyBlock.tsx";


function App() {
    const vacanciesData = [
        {
            id: 1,
            title: "Frontend Developer",
            image: "https://avatars.mds.yandex.net/i?id=dc7ff4e61803651fcae2f0fe2c5aa654cdf28846-12644621-images-thumbs&n=13",
            salary: "$3000 - $4000",
            description: "Мы ищем опытного Frontend Developer для работы над интересными проектами.",
        },
        {
            id: 2,
            title: "UX/UI Designer",
            image: "designer.jpg",
            salary: "$2500 - $3500",
            description: "Ищем талантливого UX/UI Designer с опытом работы в веб-дизайне.",
        },
        {
            id: 3,
            title: "Data Scientist",
            image: "data.jpg",
            salary: "$4000 - $5000",
            description: "Требуется профессиональный Data Scientist для анализа больших данных.",
        },



    ];

  return (
      <Router>
          <Header/>
          <NewsBlock/>
          <VacancyBlock vacancies={vacanciesData} />
          <Footer/>
      </Router>
  )
}

export default App
