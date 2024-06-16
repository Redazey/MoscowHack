import { useState } from 'react';
import './registration_style.scss';
import Header from '../components/header/Header';
import { Link } from 'react-router-dom';
import Footer_small from '../components/footer_small/footer_small';

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
            <Header />
            <div className="auth-container">
                <main className="main">
                    <div className="auth-content">
                        <div className="auth-box">
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
                        </div>
                        <div className="auth-box register-box">
                            <p>
                                Есть аккаунт? <Link to="/Login">Авторизуйтесь</Link>
                            </p>
                        </div>
                    </div>
                </main>
            </div>
            <Footer_small/>
        </>
    );
};

export default RegistrationForm;
