// @ts-ignore
import React from 'react';
import classes from './Header.module.scss';
const Header = () => {
    return (
        <header className={classes.header}>
            <div className={classes.logo}>
                <img src="../../assets/Images/HeadSoftLogo.jpg" alt="Логотип"/>
                    <h1>HeadSoft</h1>
            </div>
            <div className={classes.menu}>
                <a href="#">Главная</a>
                <a href="#">Новости</a>
                <a href="#">Вакансии</a>
            </div>
            <div className={classes.user}>
                <img src="avatar.png" alt="Аватар"/>
                    <button className={classes.login}>Войти</button>
                    <div className={classes.settings}>
                        <img src="settings.png" alt="Настройки"/>

                    </div>
            </div>
        </header>

);
};

export default Header;