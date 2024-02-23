import pandas as pd
data = pd.read_csv('datasets/olypics.csv', skiprows=4)
data
data['City']
data.City
data[['City', 'Edition', 'Sport']]
data.head()
data.head(10)
data.tail()
data.tail(15)
data.size  # no.of rows
data.shape  # tuple --> (rows, cols)

data[data['City']=='Athens']
data[data.City=='Athens']
data.City=='Athens'
data[data['Event']=='100m freestyle']

data['Sport'].value_counts()
data[data['Event']=='100m freestyle']['Athlete']  # for 1 column, just a str
data[data['Event']=='100m freestyle'][['Athlete', 'Gender', 'Medal']] # for multiple columns a list

data['Athlete'].sort_values(ascending=False)
data.sort_values(by='Edition')
data.sort_values(by='Edition', ascending=False)
data.sort_values(by=['Edition', 'Athlete'])
data.sort_values(by=['Edition', 'Athlete'], ascending=[False, True])

data[(data['Event']=='100m freestyle') & (data['Gender']=='Women')]
data[(data['Event']=='100m freestyle') & (data['Gender']=='Women') & (data['Medal']=='Gold')]
data.info
data.describe()

data[data.Athlete.str.contains('Singh')]
data[data['Athlete'].str.contains('Singh')]
data[data['Athlete'].str.startswith('BOLT')]

data['MedalCode'] = data.Medal.str.upper()

data[data['Edition'] > 2000]

groupby_city = data.groupby('City')
groupby_city['Medal'].count()  # for numerical columns, you can use min, max, sum, etc
# check the dir(groupby_city)

# import matplotlib.pyplot as plt
# %matplotlib inline

data['City'].value_counts().plot()
data['City'].value_counts().plot(kind='barh', color='#ababab')
data['Gender'].value_counts().plot(kind='pie')

