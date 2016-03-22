import json
resume = {
  'name': 'Ben Friedland',
  'email': 'resume@bugben.com',
  'url': 'http://www.bugben.com',
  'profile': [
    'Senior software/systems engineer with fifteen years of professional experience.',
    'Unique mix of software engineering and systems administration.',
  ],
  'links': [
    {'title': 'Linked-In Profile', 'url': 'http://www.linkedin.com/in/bfriedland'},
    {'title': 'GitHub Page', 'url': 'http://github.com/ben174'},
  ],
  'expertise': [
    {'category': 'Web Application Development', 'entries': ['Python', 'Django']},
    {'category': 'Mobile Development', 'entries': ['Responsive Web Design', 'REST communication (creation and consumption)']},
  ],
  'community': [
    {'title': 'HSTS Super Cookie', 'projectUrl': 'https://github.com/ben174/hsts-cookie', 'srcUrl': 'https://github.com/ben174/hsts-cookie', 'summary': 'Creates a cookie', 'body': 'Creates a super cookie'},
    {'title': 'A&G Calendar', 'projectUrl': 'http://www.armstrongandgettybingo.com/', 'srcUrl': 'https://github.com/ben174/angrates', 'summary': 'A podcast calendar for listeners to have a resource for finding episodes of my favorite radio show.', 'body': 'Listeners of the popular radio show "The Armstrong and Getty Show" were constantly having trouble finding episodes of the podcast available online. There was no clean and easy way to find episodes. I made this site as an attempt to simplify navigation and make episodes easily searchable.'},
  ],
  'experience': [
    {
        'employer': 'KIXEYE',
        'url': 'http://www.kixeye.com',
        'title': 'Lead Developer',
        'startDate': 'Creates a cookie',
        'endDate': 'Creates a cookie',
        'accomplishments': [
            'Team leader in a team of seven - in an extremely fast-paced tools environment for one of the big names in video games. Used Python + Flask to develop internal tools to manage user data within several games.',
            'Interacted with REST endpoints for several video games, creating a common UI to manage individual aspects of each game.'
        ],
        'logoUrl': 'http://yahoo.com'
    },
  ],
}
print json.dumps(resume)
