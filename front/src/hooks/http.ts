import { useCallback, useState } from 'react';
import axios, { Method, AxiosError } from 'axios';

interface IError {
    status?: number;
    message?: string;
}

export const useHttp = () => {
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState<IError | null>();

    const request = useCallback(async <T>(method: Method, url: string, data?: any): Promise<T | undefined> => {
        let response: T | undefined;
        setError(null)
        setIsLoading(true)
        try {
            response = (await axios.request({
                method,
                url,
                data,
                timeout: 5000,
            })).data
        } catch (e: any) {
            if (axios.isAxiosError(e)) {
                setError({
                    status: e.response?.status,
                    // fixme
                    // @ts-ignore
                    message: e.response?.data?.message,
                })
            } else {
                throw e
            }
        } finally {
            setIsLoading(false);
        }

        return response;
    }, []);

    return { request, isLoading , error }
}