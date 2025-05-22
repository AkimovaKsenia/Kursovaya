import { useOutside } from "hooks/useOutside";
import { FC, useEffect, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { IAuthFields } from "./login-form.interface";
import { useAuth } from "hooks/useAuth";
import styles from "./LoginForm.module.scss";
import { validEmail } from "./login-auth.constants";
import Button from "../../Button/Button";
import Field from "../../Field/Field";
import UserAvatar from "../../user-avatar/UserAvatar";
import { menuAnimation } from "utils/animation/Fade";
import { motion } from "framer-motion";
import { useMutation } from "@tanstack/react-query";
import { AuthService } from "services/auth.service";
import cn from "classnames";
import { useRouter } from "next/router";
import { TailwindLoader } from "@/components/ui/TailwindLoader";

const LoginForm: FC = () => {
  const { ref, setIsShow, isShow } = useOutside(false);
  const [type, setType] = useState<"login" | "register">("login");
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

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
      if (data?.token && data?.role) {
        setUser({
          role: data.role,
          token: data.token,
        });
        reset();
        setIsShow(false);
        console.log("✅ Пользователь вошёл", data);
      } else {
        console.warn("⚠️ Нет данных о пользователе");
        setUser(null);
      }
    },
    onError: (error: any) => {
      console.error("❌ Ошибка входа:", error?.response?.data || error.message);
    },
  });

  const onSubmit: SubmitHandler<IAuthFields> = (data) => {
    if (type == "login") loginSync.mutate(data);
    // if (type == "login")
    //   setUser({
    //     id: 1,
    //     email: "test@test.ru",
    //     name: "Dev",
    //   });
    // else if (type=='register') registerSync(data)
  };

  useEffect(() => {
    const handleRouteChange = () => setIsLoading(true);
    const handleRouteComplete = () => setIsLoading(false);
    const handleRouteError = () => setIsLoading(false);

    router.events.on("routeChangeStart", handleRouteChange);
    router.events.on("routeChangeComplete", handleRouteComplete);
    router.events.on("routeChangeError", handleRouteError);

    return () => {
      router.events.off("routeChangeStart", handleRouteChange);
      router.events.off("routeChangeComplete", handleRouteComplete);
      router.events.off("routeChangeError", handleRouteError);
    };
  }, [router]);
  useEffect(() => {
    router.prefetch("/dashboard");
  }, []);
  return (
    <div className={styles.wrapper} ref={ref}>
      {isLoading && <TailwindLoader />}
      {user ? (
        <div style={{ cursor: "pointer" }}>
          <UserAvatar link="/dashboard" title="Перейти в админ" />
        </div>
      ) : (
        <button
          style={{
            background: "none",
          }}
          className={styles.button}
          onClick={() => setIsShow(!isShow)}
        >
          <img src="/FirstUser.png" alt="User" className={styles.userIcon} />
        </button>
      )}

      <motion.div
        initial={false}
        animate={isShow ? "open" : "closed"}
        variants={menuAnimation}
      >
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
                value: 4,
                message: "Min length should more 6 symbols",
              },
            })}
          />
          <div
            className={cn(" boldText mt-1 mb-1 text-center", styles.darkText)}
          >
            <Button
              className={cn(" boldText mt-1 mb-1 text-center", styles.darkText)}
              onClick={() => setType("login")}
            >
              Login
            </Button>
          </div>
        </form>
      </motion.div>
    </div>
    //   <div className="text-black">LoginForm</div>
  );
};
export default LoginForm;
