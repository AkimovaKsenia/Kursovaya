import { FC } from "react";
import { IStatisticItem } from "./statistic-item.interface";
import styles from "./StatisticItem.module.scss";
import Image from "next/image";
const StatisticItem: FC<{ item: IStatisticItem }> = ({ item }) => {
  return (
    <div className={styles.item}>
      <item.Icon className={styles.icon} />
      <div className={styles.name}>{item.name}</div>
      <div>
        <Image
          style={{
            textAlign: "right",
          }}
          className={styles.image}
          src={item.image}
          alt={item.name}
          width={40}
          height={40}
        />
      </div>
      <div className={styles.value}>{item.value.toLocaleString("ru-RU")}</div>
    </div>
  );
};
export default StatisticItem;
