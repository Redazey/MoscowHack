// @ts-ignore
import React from 'react';
import classes from "./Footer.module.scss";
const Footer = () => {
    return (
        <footer className={classes.footer}>
            <div className={classes.column}>
                <h3>Про Reksoft</h3>
                <ul>
                    <li><a href="#">Главная страница</a></li>
                    <li><a href="#">Контакты</a></li>
                </ul>
            </div>
            <div className={classes.column}>
                <h3>О разработке</h3>
                <ul>
                    <li><a href="#">97jertov97@gmail.com</a></li>
                    <li><a href="#">1lyaaksenov@mail.com</a></li>
                </ul>
            </div>
            <div className={classes.column}>
                <h3>Колонка 3</h3>
                <ul>
                    <li><a href="#">Ссылка 1</a></li>
                    <li><a href="#">Ссылка 2</a></li>
                    <li><a href="#">Ссылка 3</a></li>
                </ul>
            </div>
        </footer>


    );
};

export default Footer;
