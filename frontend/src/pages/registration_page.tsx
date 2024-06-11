import React, { useState } from 'react';
import './registration_style.scss';

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
        // Handle form submission
    };

    return (
        <div className="registration-form">
            <h2>Регистрация</h2>
            <form onSubmit={handleSubmit}>
                <input 
                    type="text" 
                    name="firstName" 
                    placeholder="Имя" 
                    value={form.firstName} 
                    onChange={handleChange} 
                />
                <input 
                    type="text" 
                    name="lastName" 
                    placeholder="Фамилия" 
                    value={form.lastName} 
                    onChange={handleChange} 
                />
                <input 
                    type="email" 
                    name="email" 
                    placeholder="E-mail" 
                    value={form.email} 
                    onChange={handleChange} 
                />
                <input 
                    type="password" 
                    name="password" 
                    placeholder="Пароль" 
                    value={form.password} 
                    onChange={handleChange} 
                />
                <input 
                    type="password" 
                    name="confirmPassword" 
                    placeholder="Повтор пароля" 
                    value={form.confirmPassword} 
                    onChange={handleChange} 
                />
                <button type="submit">Создать аккаунт</button>
            </form>
            <div className="login-link">
                <span>Есть аккаунт?</span>
                <a href="/login">Авторизуйтесь</a>
            </div>
        </div>
    );
};

export default RegistrationForm;
