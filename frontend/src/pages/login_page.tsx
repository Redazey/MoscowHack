import React from 'react';
import ReactDOM from 'react-dom/client';
import './login_style.scss';
import Header from '../components/header/Header';
import Footer_small from '../components/footer_small/footer_small';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

const LoginForm = () => (
  <div className="auth-box">
    <h2>Авторизация</h2>
    <form>
      <input type="email" placeholder="E-mail" required />
      <input type="password" placeholder="Пароль" required />
      <button type="submit">Войти</button>
    </form>
  </div>
);

const RegisterLink = () => (
  <div className="auth-box register-box">
    <p>Нет аккаунта? <a href="#">Зарегистрируйтесь</a> </p>
  </div>
);

root.render(
  <React.StrictMode>
    <div className="auth-container">
      <Header/>
      <main className="main">
        <div className="auth-content">
          <LoginForm />
          <RegisterLink />
        </div>
      </main>
      <Footer_small/>
    </div>
  </React.StrictMode>
);
