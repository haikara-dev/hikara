import type { NextPage } from "next";
import Head from "next/head";
import { Container, Typography, Box, Stack, Card, Button } from "@mui/material";
import React, { useEffect, useState } from "react";
import AddSiteFormDialog from "@/components/site/AddSiteFormDialog";
import SiteRow from "@/components/site/SiteRow";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useAuthUserContext } from "@/lib/AuthUser";
import EditSiteFormDialog from "@/components/site/EditSiteFormDialog";
import DryRunDialog from "@/components/site/DryRunDialog";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;
const BACKEND_ADMIN_API_URL: string =
  process.env.NEXT_PUBLIC_BACKEND_ADMIN_API_URL!;

export type Site = {
  id: number;
  name: string;
  url: string;
  feed_url: string;
  active: boolean;
};

export type DryRunResult = {
  count: number;
  contents: string;
};

const Sites: NextPage = () => {
  const [data, setData] = useState<Site[]>([]);
  const [isLoading, setLoading] = useState(false);
  const { authUser } = useAuthUserContext();

  const [addOpen, setAddOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [editTarget, setEditTarget] = useState<Site | null>(null);

  const [dryOpen, setDryOpen] = useState(false);
  const [dryRunResult, setDryRunResult] = useState<DryRunResult | null>(null);

  const handleAddOpen = () => {
    setAddOpen(true);
  };

  const handleAddClose = () => {
    setAddOpen(false);
  };

  const handleEditOpen = (site: Site) => {
    setEditOpen(true);
    setEditTarget(site);
  };

  const handleEditClose = () => {
    setEditOpen(false);
    setEditTarget(null);
  };

  const openDryDialog = (result: DryRunResult) => {
    setDryRunResult(result);
    setDryOpen(true);
  };

  const handleDryClose = () => {
    setDryOpen(false);
  };

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites", {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setData(json);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const addSite = async (name: string, url: string, feed_url: string) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_ADMIN_API_URL + "/sites", {
        method: "POST",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify({
          name: name,
          url: url,
          feed_url: feed_url,
          active: true,
        }),
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const activeSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/active/"),
        {
          method: "PATCH",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            active: true,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const deActiveSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/deActive/"),
        {
          method: "PATCH",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            active: false,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const removeSite = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/"),
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

  const updateSite = async (
    id: number,
    name: string,
    url: string,
    feed_url: string,
    active: boolean
  ) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/"),
        {
          method: "PUT",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            name: name,
            url: url,
            feed_url: feed_url,
            active: active,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const runCrawling = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/run-crawling/"),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      console.log(json);
    } catch (err) {
      console.log(err);
    }
  };

  const dryRunCrawling = async (id: number) => {
    setDryRunResult(null);
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(
          id.toString(),
          BACKEND_ADMIN_API_URL + "/sites/dry-run-crawling/"
        ),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      openDryDialog(json);
    } catch (err) {
      console.log(err);
    }
  };

  const getRssUrl = async (id: number): Promise<string> => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_ADMIN_API_URL + "/sites/get-rss-url/"),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      console.log(json);
      return json.url;
    } catch (err) {
      console.log(err);
    }
    return "";
  };

  const getRssUrlByUrl = async (url: string): Promise<string> => {
    try {
      const headers = await getRequestHeaders();
      const queryParams = new URLSearchParams({ url: url });
      const res = await fetch(
        new URL(
          BACKEND_ADMIN_API_URL + "/sites/get-rss-url-by-url?" + queryParams
        ),
        {
          method: "GET",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);

      const json = await res.json();
      console.log(json);
      return json.url;
    } catch (err) {
      console.log(err);
    }
    return "";
  };
  useEffect(() => {
    setLoading(true);
    loadData();
  }, []);

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
            Sites
          </Typography>

          <Button variant="outlined" onClick={handleAddOpen}>
            Add Site
          </Button>

          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <Stack gap={2} mt={2} pr={8}>
              {data.map((site) => {
                return (
                  <Card key={site.id}>
                    <SiteRow
                      key={site.id}
                      site={site}
                      activeSite={activeSite}
                      deActiveSite={deActiveSite}
                      removeSite={removeSite}
                      updateSite={updateSite}
                      openDialog={handleEditOpen}
                      runCrawling={runCrawling}
                      dryRunCrawling={dryRunCrawling}
                    />
                  </Card>
                );
              })}
            </Stack>
          )}
        </Container>
      </Box>

      <AddSiteFormDialog
        open={addOpen}
        handleClose={handleAddClose}
        addSite={addSite}
        getRssUrlByUrl={getRssUrlByUrl}
      />

      {editTarget && (
        <EditSiteFormDialog
          open={editOpen}
          handleClose={handleEditClose}
          site={editTarget}
          updateSite={updateSite}
          onEndEdit={handleEditClose}
          getRssUrl={getRssUrl}
        />
      )}

      {dryRunResult && (
        <DryRunDialog
          open={dryOpen}
          handleClose={handleDryClose}
          dryRunResult={dryRunResult}
        />
      )}

      <Footer />
    </div>
  );
};

export default Sites;
