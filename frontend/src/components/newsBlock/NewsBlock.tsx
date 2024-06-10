import React, {useEffect, useState} from 'react';
import "./NewsBlock.module.scss";

const newsData = [
    {
        title: 'Новая коллекция осенней одежды уже в продаже',
        summary: 'Узнайте о последних трендах и сезонных скидках на осеннюю одежду.',
        image: 'https://example.com/news1.jpgааа',
    },
    {
        title: 'Завершение акции "Скидки на технику"',
        summary: 'Поторопитесь воспользоваться скидками на ноутбуки, смартфоны и другую технику!',
        image: 'https://example.com/news2.jpg',
    },
    {
        title: 'Открытие нового ресторана "Вкусные истории"',
        summary: 'Приглашаем вас на открытие нового ресторана с изысканным меню и уютной атмосферой.',
        image: 'https://example.com/news3.jpg',
    }
];


const NewsBlock: React.FC = () => {
    const [currentNewsIndex, setCurrentNewsIndex] = useState(0);

    useEffect(() => {
        const interval = setInterval(() => {
            setCurrentNewsIndex((prevIndex) => (prevIndex + 1) % newsData.length);
        }, 7000);

        return () => clearInterval(interval);
    }, []);

    const showPrevNews = () => {
        setCurrentNewsIndex((prevIndex) => (prevIndex - 1 + newsData.length) % newsData.length);
    }

    const showNextNews = () => {
        setCurrentNewsIndex((prevIndex) => (prevIndex + 1) % newsData.length);
    }

    return (
        <div className="slider">
            <div className="slider__news" style={{ backgroundImage: `url(${newsData[currentNewsIndex].image})` }}>
                <div className="slider__content">
                    <h2>{newsData[currentNewsIndex].title}</h2>
                    <p>{newsData[currentNewsIndex].summary}</p>
                    <button>Подробнее</button>
                </div>
            </div>
            <button className="slider__prev" onClick={showPrevNews}>Предыдущая новость</button>
            <button className="slider__next" onClick={showNextNews}>Следующая новость</button>
        </div>
    );
}

export default NewsBlock;