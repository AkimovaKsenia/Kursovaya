import { MotionProps, Variants } from "framer-motion";

export const FADE_IN: MotionProps = {
  initial: { opacity: 0 },
  whileInView: { opacity: 1 },
  viewport: { once: true },
  transition: {
    duration: 0.5, // Уменьшили с 1.4 до 0.5 секунды
    ease: "easeOut", // Добавили плавное завершение анимации
  },
};

export const menuAnimation: Variants = {
  open: {
    opacity: 1,
    y: 10, // Сдвигаем меню вниз относительно иконки
    filter: "blur(0px)",
    transition: {
      duration: 0.2,
      ease: "easeOut",
    },
  },
  closed: {
    opacity: 0,
    y: 10, // Начальная позиция чуть выше (для плавности)
    filter: "blur(1px)",
    transition: {
      duration: 0.15,
      ease: "easeIn",
    },
  },
};
