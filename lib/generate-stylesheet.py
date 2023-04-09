import re, yaml

def append_initial_css():
    return """
.currency-flag {
  display: inline-block;
  width: 24px;
  height: 16px;
  background-size: cover;
}
.currency-flag-sm {
  width: 16px;
  height: 10px;
}
.currency-flag-lg {
  width: 36px;
  height: 24px;
}
.currency-flag-xl {
  width: 48px;
  height: 32px;
}
"""

def make_css_class(code, flag):
    return f"""
    .currency-flag.currency-flag-{code.lower()} {{
      background-image: url({flag});
    }}
    """

def minify_css(css):
    return re.sub(r"\s+", " ", css)

def __main__():
  # Get initial CSS, ready to combine with generated styles
  stylesheet_content = append_initial_css()

  # Read file listing supported currencies
  with open('../data/currencies.yaml') as file:
      try:
          currencies = yaml.safe_load(file) or []
          # For each currency, generate CSS class, and append
          for item in currencies:
            stylesheet_content += make_css_class(item['code'], item['flag'])
          
      except yaml.YAMLError as exc:
          print(exc)

  # Write generated CSS to file
  with open("../css/currency-flags.css", "w") as file:
      file.write(stylesheet_content)

  # Minify styles, and write to min.css file
  with open("../css/currency-flags.min.css", "w") as file:
      file.write(minify_css(stylesheet_content))

