import React from 'react';
import './VacanciesBlock.scss';
interface Vacancy {
    image: string;
    title: string;
    salary: string;
    description: string;
}

interface VacancyBlockProps {
    vacancies: Vacancy[];
}

const VacancyBlock: React.FC<VacancyBlockProps> = ({ vacancies }) => {
    return (
        <div className="vacancies-block">
            {vacancies.map((vacancy, index) => (
                <div key={index} className="vacancy-card">
                    <img src={vacancy.image} alt={vacancy.title} />
                    <h3>{vacancy.title}</h3>
                    <p>Salary: {vacancy.salary}</p>
                    <p>{vacancy.description}</p>
                    <button>Откликнуться</button>
                </div>
            ))}
        </div>
    );
}

export default VacancyBlock;
