
import styles from 'components/ClubLeague/loadBar.module.scss';
import { ReactNode } from 'react';

interface DayProps {
    child?: ReactNode;
    cnt: number;
    isActive: boolean;
}

const Day = ({ isActive, cnt }: DayProps) => {
    return (
        <div className={styles.loadBarDay}>
            <div className={`${styles.loadBarDayBall} ${isActive ? styles.loadBarDayBallActive : ""}`}/>
            <span className={styles.loadBarDayTitle}>day {cnt}</span>
        </div>
    )
}

export const LoadBar = () => {
    return (
        <div className={styles.loadBar}>
            <div className={styles.loadBarDays}>
                {[1, 2, 3, 4, 5].map(idx => <Day isActive={idx < 3} cnt={idx}/>)}
            </div>
        </div>
    );
}