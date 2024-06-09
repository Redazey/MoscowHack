// @ts-ignore
import React,  { useState } from 'react';
import classes from './Header.module.scss';
import Logo from '../../assets/Images/HeadSoftLogo.jpg';
import Profile from '../../assets/Images/Profile.jpg';

const Header = () => {
    const [isOpen, setIsOpen] = useState(false);

    const toggleDropdown = () => {
        setIsOpen(!isOpen);
    };

    return (
        <nav className={classes.header}>
            <div className={classes.logo}>
                <img  id="scrolling-image"  src={Logo} alt="Логотип"/>
                    <h1>HeadSoft</h1>
            </div>
            <div className={classes.menu}>
                <a href="#">Главная</a>
                <a href="#">Новости</a>
                <a href="#">Вакансии</a>
            </div>
            <div className={classes.user}>
                <img src={Profile} alt="Аватар"/>
                <a><div className={classes.login}>Войти</div></a>
                <button onClick={toggleDropdown} className="dropdown-btn">
                    Меню
                </button>
                {isOpen && (
                    <div className="dropdown-menu">
                        <ul>
                            <li>Пункт 1</li>
                            <li>Пункт 2</li>
                            <li>Пункт 3</li>
                        </ul>
                    </div>
                )}
            </div>
        </nav>

);
};

export default Header;