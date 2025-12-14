CREATE TABLE IF NOT EXISTS recipes (
  id      INT AUTO_INCREMENT NOT NULL,
  title   VARCHAR(128) NOT NULL,
  slug    VARCHAR(128) NOT NULL,
  credit  VARCHAR(128),
  image   VARCHAR(255),
  portions  JSON,
  ingredients JSON,
  steps JSON,
  PRIMARY KEY (`id`)
);

INSERT INTO recipes
  (title, slug, credit, image, portions, ingredients, steps)
VALUES 
  ('Sourdough', 'sourdough', '', '', '{"unit": "loaf", "units": "loaves", "value": 1}', '[  ]', '[  ]');
