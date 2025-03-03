import React, { useEffect, useState } from 'react';
import './loading.css';

interface LoadingProps {
    size?: 'small' | 'medium' | 'large';
}

// 预加载图片
const preloadImage = (src: string): Promise<void> => {
    return new Promise((resolve, reject) => {
        const img = new Image();
        img.onload = () => resolve();
        img.onerror = reject;
        img.src = src;
    });
};

const LOGO_URL = 'https://shifu.dev/img/logo.svg';

// 立即开始预加载
preloadImage(LOGO_URL).catch(console.error);

const Loading: React.FC<LoadingProps> = ({ size = 'medium' }) => {
    const [isAnimating, setIsAnimating] = useState(false);
    const [imageLoaded, setImageLoaded] = useState(false);

    useEffect(() => {
        // 确保图片已加载
        preloadImage(LOGO_URL)
            .then(() => {
                setImageLoaded(true);
                setIsAnimating(true);
            })
            .catch(console.error);

        return () => setIsAnimating(false);
    }, []);

    const sizeClass = {
        small: 'w-8 h-8',
        medium: 'w-16 h-16',
        large: 'w-24 h-24'
    }[size];

    if (!imageLoaded) {
        return <div className="loading-container" />;  // 占位
    }

    return (
        <div className="loading-container">
            <img
                src={LOGO_URL}
                alt="Loading..."
                className={`loading-spinner ${sizeClass} rounded-full ${isAnimating ? 'animate' : ''}`}
            />
            <div className={`loading-pulse ${sizeClass} bg-blue-500 ${isAnimating ? 'animate' : ''}`}></div>
        </div>
    );
};

export default Loading;
