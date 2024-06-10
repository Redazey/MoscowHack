import React, {useEffect, useState} from 'react';
import classes from './NewsBlock.module.scss';
import NewsImage from '../../assets/Images/NewsImage.jpg';
import { motion } from 'framer-motion';

const newsData = [
    {
        title: 'Новость 1',
        summary: 'Описание Новости 1 маленькое смешное',
        image: '../../assets/Images/settings-cog.jpg',
    },
    {
        title: 'Новость 2',
        summary: 'Описание Новости 1 маленькое смешное',
        image: '../../assets/Images/settings-cog.jpg',
    },
    {
        title: 'Новость 3',
        summary: 'Описание Новости 1 маленькое смешное',
        image: '../../assets/Images/settings-cog.jpg',
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
        <div className={classes.slider}>
            <motion.div
                className={classes.slider__news}
                style={{backgroundImage: `url(${newsData[currentNewsIndex].image})`}}
                initial={{opacity: 0}}
                animate={{opacity: 1}}
                transition={{duration: 0.5}}>
                <div className={classes.slider__content}>
                    <h2>{newsData[currentNewsIndex].title}</h2>
                    <p>{newsData[currentNewsIndex].summary}</p>
                    <button>Подробнее</button>
                </div>
            </motion.div>
            <button className={classes.slider__prev} onClick={showPrevNews}>←</button>
            <button className={classes.slider__next} onClick={showNextNews}>→</button>
        </div>
    );
}

export default NewsBlock;