import React from "react";

import {
  Container,
  FileCard,
  FilesGrid,
  FileOptions,
  Layout,
  SearchBar,
} from "../../components";

const Home: React.FC = () => {
  return (
    <Layout>
      <Container>
        <SearchBar />
        <FileOptions />
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
          <FileCard name="file_name.png" />
          <FileCard name="file_name.png" />
          <FileCard name="file_name.png" />
          <FileCard name="file_name.png" />
        </FilesGrid>
      </Container>
    </Layout>
  );
};

export default Home;
