import React from "react";

import { Container, FileCard, FilesGrid, Layout } from "../../components";

const Home: React.FC = () => {
  return (
    <Layout>
      <Container>
        <FilesGrid>
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
          <FileCard name="file_name.png" />
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
          <FileCard name="file_name.png" />
          <FileCard
            name="file_name.png"
            thumbnail="https://c.wallhere.com/photos/67/6e/vertical_portrait_display-1183234.jpg!d"
          />
        </FilesGrid>
      </Container>
    </Layout>
  );
};

export default Home;
