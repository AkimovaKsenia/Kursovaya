import Layout from "@/components/ui/layout/Layout";
import { FC, useEffect, useRef } from "react";
import Link from "next/link";
import Button from "../../components/ui/layout/Button/Button";
import { useAuth } from "hooks/useAuth";
import { useRouter } from "next/router";
import { AuthService } from "services/auth.service";
import styles from "./Home.module.scss";
import cn from "classnames";
import gsap from "gsap";

const Home: FC = () => {
  const { user, setUser } = useAuth();
  const router = useRouter();
  const animationRef = useRef<gsap.Context | null>(null);
  const contentRef = useRef<HTMLDivElement>(null);

  const handleLogout = () => {
    AuthService.logout();
    setUser(null);
  };
  console.log("Rendering Home page, user:", user);
  console.log("Rendering Home page");

  useEffect(() => {
    if (animationRef.current) {
      animationRef.current.revert();
    }

    animationRef.current = gsap.context(() => {
      gsap.from(contentRef.current, {
        duration: 0.8,
        opacity: 0,
        y: 30,
        ease: "power2.out",
        delay: 0.3,
      });
    });
    return () => {
      if (animationRef.current) {
        animationRef.current.revert();
      }
    };
  }, [router.pathname]);

  return (
    <Layout
      title="Cinema"
      backgroundImage="/imageMain.png"
      backgroundColor="#1E1E1E"
    >
      <div
        ref={contentRef}
        style={{
          backgroundBlendMode: "overlay",
        }}
        className="flex flex-col items-center mr-210 justify-center min-h-[60vh] space-y-2 mt-21"
      >
        <h1 className={styles.heading}>Cinema</h1>
        <h1 className={styles.headingMastery}>MASTERY</h1>
        <div className="max-w-110">
          <h3 className={styles.text}>
            Войдите в систему и просматривайте актуальную информацию о сеансах,
            фильмах и кинотеатрах города!
          </h3>
        </div>
        {user && (
          <div className="flex items-center space-x-4">
            <div className="flex space-x-4 mt-6">
              <Link href="/ChoiceCinema" passHref>
                <Button
                  className={cn(
                    "px-6 py-3 text-lg mt-0 boldText",
                    styles.darkText
                  )}
                >
                  Кинотеатры
                </Button>
              </Link>
              <Link href="/ChoiceMovie" passHref>
                <Button
                  className={cn(
                    "px-6 py-3 text-lg mt-0 boldText",
                    styles.darkText
                  )}
                >
                  Фильмы
                </Button>
              </Link>
            </div>
          </div>
        )}
        <div>
          <button
            onClick={handleLogout}
            style={{ cursor: "pointer" }}
            className={cn("px-6 py-3 text-lg mt-0 boldText")}
          >
            Выйти
          </button>
        </div>{" "}
      </div>
      <div className="absolute mt-160 inset-0 bg-gradient-to-t from-black to-transparent" />
      <div className="absolute mt-160 bottom-8 left-8 right-8"></div>
    </Layout>
  );
};
export default Home;
