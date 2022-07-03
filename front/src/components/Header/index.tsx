import { ReactNode } from 'react';

import styles from 'components/Header/header.module.scss';

interface IHeaderProps {
    child?: ReactNode;
    tag?: string;
    name?: string;
    trophies?: number;
}

export const Header = (props: IHeaderProps) => {
    return (
        <header className={styles.header}>
            <span className={styles.headerTag}>[{props.tag}]</span>
            <span className={styles.headerName}>{props.name}</span>&nbsp;
            <span className={styles.headerTrophies}>{props.trophies}🏆</span>
        </header>
    );
}
