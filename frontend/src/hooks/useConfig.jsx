import { useCallback, useEffect } from 'react';
import { useGetState } from './useGetState';

export const useConfig = (key, defaultValue, options = {}) => {
    const [property, setPropertyState, getProperty] = useGetState(null);

    // 同步到State (Store -> State)
    const syncToState = useCallback((v) => {
        if (v !== null) {
            setPropertyState(v);
        } else {
            if (getProperty() === null) {
                setPropertyState(defaultValue);
            } else {
                setPropertyState(v);
            }
        }
    }, []);

    const setProperty = useCallback((v, forceSync = false) => {
        setPropertyState(v);
    }, []);

    // 初始化
    useEffect(() => {
        syncToState(null);
        if (key.includes('[')) return;
    }, []);

    return [property, setProperty, getProperty];
};
