import React from 'react';
import './loading.css';

interface LoadingProps {
    size?: 'small' | 'medium' | 'large';
}

const Loading: React.FC<LoadingProps> = ({ size = 'medium' }) => {
    const sizeClass = {
        small: 'w-8 h-8',
        medium: 'w-16 h-16',
        large: 'w-24 h-24'
    }[size];

    return (
        <div className="loading-container">
            <img
                src="https://shifu.dev/img/logo.svg"
                alt="Loading..."
                className={`loading-spinner ${sizeClass}`}
            />
            <div className={`loading-pulse ${sizeClass}`}></div>
        </div>
    );
};

export default Loading;
