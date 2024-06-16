import classes from './Header.module.scss';
import Logo from '../../assets/Images/HeadSoftLogo.jpg';
import Profile from '../../assets/Images/Profile.jpg';
import SettingsCog from '../../assets/Images/settings-cog.jpg';
import { Link } from 'react-router-dom';

const Header = () => {
    return (
        <nav className={classes.header}>
            <div className={classes.logo}>
                <img  id="scrolling-image"  src={Logo} alt="Логотип"/>
                <h1>HeadSoft</h1>
            </div>
            <div className={classes.menu}>
                <Link to="/">Главная</Link>
                <a href="#">Новости</a>
                <a href="#">Вакансии</a>
            </div>
            <div className={classes.user}>
                <img src={Profile} alt="Аватар"/>
                <Link to="/Login" className={classes.login}>Войти</Link>
                <img src={SettingsCog} className="dropdown-btn"/>
            </div>
        </nav>
    );
};

export default Header;