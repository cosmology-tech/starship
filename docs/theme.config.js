import { useConfig } from "nextra-theme-docs";
import React from "react";
import { useRouter } from "next/router";

/* eslint sort-keys: error */
/**
 * @type {import('nextra-theme-docs').DocsThemeConfig}
 */
export default {
  footer: {
    // text: `BSD ${new Date().getFullYear()} © ComosKit.`,
    text: (
      <span>
        🛠 Built by Cosmology — if you like our tools, please consider delegating
        to{" "}
        <a
          href="https://cosmology.tech/validator"
          target="_blank"
          rel="noreferrer"
          aria-selected="false"
          className="nx-text-primary-500 nx-underline nx-decoration-from-font [text-underline-position:under]"
        >
          our validator ⚛️
        </a>
      </span>
    ),
  },
  head: () => {
    const { asPath, defaultLocale, locale } = useRouter();
    const { title } = useConfig();
    const url =
        "https://starship.cosmology.tech/" +
        (defaultLocale === locale ? asPath : `/${locale}${asPath}`);

    const _title = asPath !== "/" ? `${title} - Cosmology` : `${title}`;
    return (
        <>
          <meta property="og:url" content={url} />
          <meta property="og:title" content={_title} />
          <meta
              property="og:description"
              content={"Unified development environment"}
          />
          <title>{_title}</title>
        </>
    );
  },
  chat: {
    link: "https://discord.gg/6hy8KQ9aJY",
  },
  project: {
    link: "https://github.com/hyperweb-io/starship",
  },
  docsRepositoryBase:
    "https://github.com/hyperweb-io/starship/tree/main/docs",
  editLink: {
    text: "Edit this page on GitHub",
  },
  getNextSeoProps() {
    const { frontMatter } = useConfig();
    return {
      additionalLinkTags: [
        {
          href: "/apple-icon-180x180.png",
          rel: "apple-touch-icon",
          sizes: "180x180",
        },
        {
          href: "/android-icon-192x192.png",
          rel: "icon",
          sizes: "192x192",
          type: "image/png",
        },
        {
          href: "/favicon-96x96.png",
          rel: "icon",
          sizes: "96x96",
          type: "image/png",
        },
        {
          href: "/favicon-32x32.png",
          rel: "icon",
          sizes: "32x32",
          type: "image/png",
        },
        {
          href: "/favicon-16x16.png",
          rel: "icon",
          sizes: "16x16",
          type: "image/png",
        },
      ],
      additionalMetaTags: [
        { content: "en", httpEquiv: "Content-Language" },
        { content: "CosmosKit", name: "apple-mobile-web-app-title" },
        { content: "#fff", name: "msapplication-TileColor" },
        { content: "/ms-icon-144x144.png", name: "msapplication-TileImage" },
      ],
      description: "CosmosKit: A wallet connector for the Cosmos ",
      openGraph: {
        images: [
          {
            url: "https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png",
          },
        ],
      },
      titleTemplate: "%s – CosmosKit",
      twitter: {
        cardType: "summary_large_image",
        site: "https://cosmoskit.com/",
      },
    };
  },
  logo: (
    <>
      <img
        src="https://user-images.githubusercontent.com/10805402/242348990-c141d6cd-e1c9-413f-af68-283de029c3a4.png"
        width="50px"
      ></img>
      <span className="mr-2 font-extrabold hidden md:inline">
        &nbsp;&nbsp;Starship
      </span>
    </>
  ),
};
