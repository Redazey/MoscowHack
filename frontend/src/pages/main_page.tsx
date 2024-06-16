import React from 'react';
import './main_style.scss';
import Header from '../components/header/Header.tsx';
import Footer from "../components/footer/Footer.tsx";
import NewsBlock from "../components/newsBlock/NewsBlock.tsx";
import VacancyBlock from "../components/VacanciesBlock/VacancyBlock.tsx";

const MainPage: React.FC = () => {
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
            image: "https://avatars.mds.yandex.net/i?id=7b54c931dac60a19dbcfa251f839cde0a2536a4d-9221733-images-thumbs&n=13",
            salary: "$2500 - $3500",
            description: "Ищем талантливого UX/UI Designer с опытом работы в веб-дизайне.",
        },
        {
            id: 3,
            title: "Data Scientist",
            image: "https://avatars.mds.yandex.net/i?id=3d1dc9fe8e02b6a1f20e4feae2f158a32158cdf2-9100256-images-thumbs&n=13",
            salary: "$4000 - $5000",
            description: "Требуется профессиональный Data Scientist для анализа больших данных.",
        },
    ];

    return (
        <>
            <Header/>
            <NewsBlock/>
            <VacancyBlock vacancies={vacanciesData} />
            <Footer/>
        </>
    )
}

export default MainPage;