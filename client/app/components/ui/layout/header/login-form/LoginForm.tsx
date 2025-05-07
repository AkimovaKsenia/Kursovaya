import { useOutside } from "hooks/useOutside";
import { FC, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import Meta from "utils/Meta";
import { IAuthFields } from "./login-form.interface";
import { useAuth } from "hooks/useAuth";
import styles from "./LoginForm.module.scss";
import { FaUserCircle } from "react-icons/fa";
import { validEmail } from "./login-auth.constants";
import Button from "../../Button/Button";
import Field from "../../Field/Field";
import UserAvatar from "../../user-avatar/UserAvatar";
import { menuAnimation } from "utils/animation/Fade";
import { motion } from "framer-motion";
import { useMutation } from "@tanstack/react-query";
import { AuthService } from "services/auth.service";

const LoginForm: FC = () => {
  const { ref, setIsShow, isShow } = useOutside(false);
  const [type, setType] = useState<"login" | "register">("login");
  const {
    register,
    formState: { errors },
    handleSubmit,
    reset,
  } = useForm<IAuthFields>({
    mode: "onChange",
  });

  const { user, setUser } = useAuth();

  const loginSync = useMutation({
    mutationKey: ["login"],
    mutationFn: (data: IAuthFields) =>
      AuthService.login(data.email, data.password),
    onSuccess: (data) => {
      if (setUser) setUser(data.user);
      reset();
      setIsShow(false);
    },
  });

  const onSubmit: SubmitHandler<IAuthFields> = (data) => {
    if (type == "login") loginSync.mutate(data);
    // else if (type=='register') registerSync(data)
  };
  return (
    <div className={styles.wrapper} ref={ref}>
      {user ? (
        <UserAvatar link="/dashboard" title="Перейти в админ" />
      ) : (
        <button
          style={{
            background: "none",
          }}
          className={styles.button}
          onClick={() => setIsShow(!isShow)}
        >
          <FaUserCircle fill="#A4A4A4" className={styles.userIcon} />
        </button>
      )}

      <motion.div animate={isShow ? "open" : "closed"} variants={menuAnimation}>
        <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
          <Field
            className={styles.input}
            type="email"
            placeholder="Email"
            error={errors.email}
            {...register("email", {
              required: "Email is required",
              pattern: {
                value: validEmail,
                message: "Please enter a valid email address",
              },
            })}
          />
          <Field
            className={styles.input}
            type="password"
            placeholder="Пароль"
            error={errors.password}
            {...register("password", {
              required: "password is required",
              minLength: {
                value: 6,
                message: "Min length should more 6 symbols",
              },
            })}
          />
          <div className={"mt-5 mb-1 text-center"}>
            <Button onClick={() => setType("login")}>Login</Button>
          </div>
          <button
            className={styles.register}
            onClick={() => setType("register")}
          >
            Register
          </button>
        </form>
      </motion.div>
    </div>
    //   <div className="text-black">LoginForm</div>
  );
};
export default LoginForm;
