import { useEffect, useState } from 'react';

import { Header } from 'components/Header';
import { Footer } from 'components/Footer';
import { Members } from 'components/Members';

import { useHttp } from 'hooks/http';

import { IClub } from 'types';
import { Description } from 'components/Description';
import { JoinRequestForm } from 'components/Form';
import { Loader } from 'components/Loader';

export const MainLayout = () => {
    const [club, setClub] = useState<IClub | undefined>();

    const {request, isLoading, error} = useHttp();

    useEffect(() => {
        request<IClub>('get', '/api/v1/clubs').then(club => {
            setClub(club);
        });
    }, [request, setClub]);

    if (isLoading) {
        return <Loader />;
    }

    return (
        <>
            <Header name={club?.name} tag={club?.tag} trophies={club?.trophies}/>
            <main>
                <JoinRequestForm />
                { club?.description && <Description description={club?.description} /> }
                { club?.members && <Members members={club?.members}/> }
            </main>
            <Footer/>
        </>
    );
};
