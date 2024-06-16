import { useState } from 'react';
import './registration_style.scss';
import Header from '../components/header/Header';
import Footer_small from '../components/footer_small/footer_small';
import { Link } from 'react-router-dom';

const RegistrationForm: React.FC = () => {
    const [form, setForm] = useState({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setForm({ ...form, [name]: value });
    };

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
    };

    return (
        <>
            <Header/>
            <div className="registration-form">
                <h2>Регистрация</h2>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        name="firstName"
                        placeholder="Имя"
                        value={form.firstName}
                        onChange={handleChange} />
                    <input
                        type="text"
                        name="lastName"
                        placeholder="Фамилия"
                        value={form.lastName}
                        onChange={handleChange} />
                    <input
                        type="email"
                        name="email"
                        placeholder="E-mail"
                        value={form.email}
                        onChange={handleChange} />
                    <input
                        type="password"
                        name="password"
                        placeholder="Пароль"
                        value={form.password}
                        onChange={handleChange} />
                    <input
                        type="password"
                        name="confirmPassword"
                        placeholder="Повтор пароля"
                        value={form.confirmPassword}
                        onChange={handleChange} />
                    <button type="submit">Создать аккаунт</button>
                </form>
                <div className="login-link">
                    <span>Есть аккаунт?</span>
                    <Link to="/Login">Авторизуйтесь</Link>
                </div>
            </div>
            <Footer_small/>
        </>
    );
};

export default RegistrationForm;
