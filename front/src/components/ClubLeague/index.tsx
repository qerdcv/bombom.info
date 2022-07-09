import styles from 'components/ClubLeague/club-league.module.scss';
import { LoadBar } from 'components/ClubLeague/loadBar';

export const ClubLeague = () => {
    return (
        <div className={styles.clubLeague}>
            <h3>Club League</h3>
            <LoadBar />
        </div>
    );
}
