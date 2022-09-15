import type { NextPage } from "next";
import Head from "next/head";
import {
  Container,
  Typography,
  Box,
  Stack,
  Card,
  Button,
  IconButton,
  Pagination,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useAuthUserContext } from "@/lib/AuthUser";
import DeleteIcon from "@mui/icons-material/Delete";
import { useRouter } from "next/router";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export type Article = {
  id: number;
  title: string;
  published_at: string;
};

const Articles: NextPage = () => {
  const router = useRouter();
  const [page, setPage] = useState<number>(
    router.query.page ? parseInt(router.query.page.toString()) : 1
  );
  const [totalCount, setTotalCount] = useState<number>(0);
  const [totalPage, setTotalPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);
  const [data, setData] = useState<Article[]>([]);
  const [isLoading, setLoading] = useState(false);
  const { authUser } = useAuthUserContext();

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const queryParams = new URLSearchParams({ page: page.toString() });
      const res = await fetch(BACKEND_API_URL + "/articles?" + queryParams, {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setTotalCount(json.totalCount);
      setTotalPage(json.totalPage);
      setPageSize(json.pageSize);
      setData(json.data);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const removeArticle = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/articles/"),
        {
          method: "DELETE",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const onClickDeleteHandler = (
    id: number,
    e: React.MouseEvent<HTMLButtonElement>
  ) => {
    e.preventDefault();
    removeArticle(id);
  };

  const handleChangePagination = (
    e: React.ChangeEvent<unknown>,
    page: number
  ) => {
    router.push({ query: { page: page } });
  };

  useEffect(() => {
    setLoading(true);
  }, []);

  useEffect(() => {
    setPage(router.query.page ? parseInt(router.query.page.toString()) : 1);
  }, [router]);

  useEffect(() => {
    loadData();
  }, [page]);

  return (
    <div>
      <Head>
        <title>DailyFJ</title>
        <meta name="description" content="DailyFJ" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <Box
        component="main"
        sx={{
          minHeight: "100vh",
        }}
      >
        <Container
          sx={{
            p: 2,
          }}
        >
          <Typography variant="h3" component="h1">
            Articles
          </Typography>

          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <Stack gap={3} alignItems="center">
              <Stack>
                {totalCount}件中　{(page - 1) * pageSize + 1} -{" "}
                {(page - 1) * pageSize + data.length}件
              </Stack>
              <Stack gap={2} mt={2} pr={8}>
                {data.map((article) => {
                  return (
                    <Card key={article.id}>
                      <Stack direction="row" gap={3} alignItems="center">
                        <div>
                          {new Date(article.published_at).toLocaleDateString()}
                        </div>
                        <Box
                          sx={{
                            flexGrow: 1,
                          }}
                        >
                          {article.title}
                        </Box>
                        <IconButton
                          onClick={onClickDeleteHandler.bind(this, article.id)}
                          aria-label="remove"
                        >
                          <DeleteIcon />
                        </IconButton>
                      </Stack>
                    </Card>
                  );
                })}
              </Stack>
              <Pagination
                page={page}
                count={totalPage}
                onChange={handleChangePagination}
              />
            </Stack>
          )}
        </Container>
      </Box>

      <Footer />
    </div>
  );
};

export default Articles;
