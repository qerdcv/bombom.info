import { ReactNode } from 'react';

import styles from 'components/Description/description.module.scss';


interface IDescriptionProps {
    child?: ReactNode;
    description: string;
}

export const Description = ({description}: IDescriptionProps) => {
    return (
        <div className={styles.description}>
            {description}
        </div>
    )
};
