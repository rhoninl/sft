import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { RouterProvider, } from "react-router-dom";
import { router } from './router';
import { HeroUIProvider, ToastProvider } from "@heroui/react";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import { Provider } from 'react-redux';
import { store } from './store/store';
import EasterEgg from './components/easteregg/EasterEgg';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <NextThemesProvider defaultTheme="light" attribute="class" enableSystem>
        <HeroUIProvider>
          <ToastProvider placement='bottom-left' />
          <main className="min-h-screen text-foreground bg-background">
            <EasterEgg />
            <div className='h-full w-full flex justify-center'>
              <RouterProvider router={router} />
            </div>
          </main>
        </HeroUIProvider>
      </NextThemesProvider>
    </Provider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
