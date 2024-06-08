// @ts-ignore
import React from 'react';
import classes from "./Footer.module.scss";
const Footer = () => {
    return (
        <footer className={classes.footer}>
            <div className={classes.columnLeft}>
                <h3>Про Reksoft</h3>
                <ul>
                    <li><a href="#">Главная страница</a></li>
                    <li><a href="#">Контакты</a></li>
                </ul>
            </div>
            <div className={classes.columnRight}>
                <h3>Связатся с нами</h3>
                <ul>
                    <li><a href="#">97jertov97@gmail.com</a></li>
                    <li><a href="#">1lyaaksenov@mail.com</a></li>
                </ul>
            </div>
            <div className={classes.columnRight}>
                <h3>О разработке</h3>
                <ul>
                    <li>
                       <a>Сделанно командой малосольные огурчики в рамках Хакатона: "Лидеры цифровой трансформации"</a>
                    </li>
                </ul>
            </div>
        </footer>


    );
};

export default Footer;
