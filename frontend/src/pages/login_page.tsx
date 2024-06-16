import React from 'react';
import './login_style.scss';
import Header from '../components/header/Header';
import Footer_small from '../components/footer_small/footer_small';
import { Link } from 'react-router-dom';

const Login: React.FC = () => {
    return (
        <>
            <Header/>
                <div className="auth-container">
                    <main className="main">
                        <div className="auth-content">
                            <div className="auth-box">
                                <h2>Авторизация</h2>
                                <form>
                                <input type="email" placeholder="E-mail" required />
                                <input type="password" placeholder="Пароль" required />
                                <button type="submit">Войти</button>
                                </form>
                            </div>
                            <div className="auth-box register-box">
                                <p>
                                    Нет аккаунта? <Link to="/Registration">Зарегистрируйтесь</Link> 
                                </p>
                            </div>
                        </div>
                    </main>
                </div> 
            <Footer_small/>
        </>
    )
}

export default Login;